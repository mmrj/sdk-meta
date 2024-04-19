import * as fsp from 'node:fs/promises';
import {
  LogCodes,
  System,
  Class,
  Code,
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
    write(`${text}\n`);
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

  async function writeCodeDocComment(definition: Code) {
    await withDocComment(async () => {
      await writeCommentLn(definition.description);
    });
  }

  async function writeMessageFunctionDocComment(definition: Code) {
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

  for (let systemName of Object.keys(definitions)) {
    await withDocComment(async () => {
      await writeCommentLn("Standardized log codes and messages.");
    });
    await scoped('export class LogMessages {', '}', async () => {
      await writeSystem(systemName, definitions[systemName]);
    });
  }

  async function writeClass(className: string, system: System, cls: Class) {
    await withDocComment(async () => {
      await writeCommentLn(cls.description);
    });
    await scoped(`static ${capitalize(className)} = class {`, '}', async () => {
      for (let codeName of Object.keys(cls.codes)) {
        const definition: Code = cls.codes[codeName];
        await writeCode(codeName, system, cls, definition);
      }
    });
  }

  async function writeSystem(systemName: string, system: System) {
    await withDocComment(async () => {
      await writeCommentLn(system.description);
    });
    await scoped(`static ${capitalize(systemName)}  = class {`, '}', async () => {
      for (let className of Object.keys(system.classes)) {
        const cls: Class = system.classes[className];
        await writeClass(className, system, cls);
      }
    });
  }

  async function writeCode(codeName: string, system: System, cls: Class, Code: Code) {
    await writeCodeDocComment(Code);
    await scoped(`static ${capitalize(codeName)}  = class {`, '}', async () => {
      const code = `${system.specifier}:${cls.specifier}:${Code.specifier}`;
      await writeLn(`static readonly code = \"${code}\";`);
      await writeMessageFunctionDocComment(Code);
      await scoped(`static message(${makeParams(Code.message)}): string {`, '}', async () => {
        await writeLn(`return \`${code} ${Code.message.parameterized}\`;`);
      });
    });
  }
}

main();
