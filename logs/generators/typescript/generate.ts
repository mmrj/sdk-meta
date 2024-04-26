import * as fsp from 'node:fs/promises';
import {
  LogCodes,
  System,
  Class,
  Condition,
  Message
} from './ld_log_codes';

function capitalize(input: string): string {
  return input.charAt(0).toUpperCase() + input.slice(1);
}

function makeParams(message: Message): string {
  return Object.keys(message.parameters || {}).map(param => `${param}: string`).join(", ");
}

async function main() {
  const definitions: LogCodes = JSON.parse(await fsp.readFile("../../data/codes.json", { encoding: "utf8" }))

  const outputFile = await fsp.open("codes.ts", "w");
  let depth = 0;

  async function write(text: string) {
    await fsp.appendFile(outputFile, `${'  '.repeat(depth)}${text}`);
  }

  async function writeLn(text: string) {
    await write(`${text}\n`);
  }

  async function scoped(start: string, end: string, scope: () => Promise<void>): Promise<void> {
    await writeLn(start);
    depth++;
    await scope();
    depth--;
    await writeLn(end);
  }

  async function withDocComment(makeComments: () => Promise<void>) {
    await writeLn("/**");
    await makeComments();
    await writeLn("*/");
  }

  async function writeCommentLn(comment: string) {
    await writeLn(` * ${comment}`);
  }

  async function writeCommentParamLn(name: string, def: string) {
    await writeCommentLn(`@param ${name} ${def}`);
  }

  async function writeConditionDocComment(definition: Condition) {
    await withDocComment(async () => {
      await writeCommentLn(definition.description);
    });
  }

  async function writeMessageFunctionDocComment(definition: Condition) {
    await withDocComment(async () => {
      await writeCommentLn("Generate a log string for this code.");
      await writeCommentLn("");
      await writeCommentLn("This function will automatically include the log code.");

      for (let paramName of Object.keys(definition.message.parameters || {})) {
        const paramDef = definition.message.parameters![paramName];
        await writeCommentParamLn(paramName, paramDef);
      }
    });
  }

  await writeLn("// This code is automatically generated and should not be manually edited.");
  await writeLn("");

  await withDocComment(async () => {
    await writeCommentLn("Standardized log codes and messages.");
  });
  await scoped('export class LogMessages {', '}', async () => {
    for (let systemName of Object.keys(definitions.systems || {})) {
      await writeSystem(systemName, definitions.systems![systemName]);
    }
  });

  async function writeClass(className: string, cls: Class, systemName: string, system: System) {
    await withDocComment(async () => {
      await writeCommentLn(cls.description);
    });
    await scoped(`static ${capitalize(className)} = class {`, '}', async () => {
      const applicable = Object.entries(definitions.conditions).filter(([_key, value]) => {
        return value.class == cls.specifier && value.system == system.specifier;
      });
      for (let [conditionCode, condition] of applicable) {
        await writeCondition(conditionCode, systemName, className, condition);
      }
    });
  }

  async function writeSystem(systemName: string, system: System) {
    await withDocComment(async () => {
      await writeCommentLn(system.description);
    });
    await scoped(`static ${capitalize(systemName)}  = class {`, '}', async () => {
      for (let [className, cls] of Object.entries(definitions.classes)) {
        await writeClass(className, cls, systemName, system);
      }
    });
  }

  async function writeCondition(conditionCode: string, system: string, cls: string, condition: Condition) {
    await writeConditionDocComment(condition);
    await scoped(`static ${capitalize(condition.name)} = class {`, '}', async () => {
      await writeLn(`static readonly code = \"${conditionCode}\";`);
      await writeMessageFunctionDocComment(condition);
      await scoped(`static message(${makeParams(condition.message)}): string {`, '}', async () => {
        await writeLn(`return \`${conditionCode} ${condition.message.parameterized}\`;`);
      });
    });
  }
}

main();
