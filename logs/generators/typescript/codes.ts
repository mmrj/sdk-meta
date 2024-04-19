// This code is automatically generated and should not be manually edited.

/**
 * Standardized log codes and messages.
*/
export class LogMessages {
  /**
   * Conditions associated with SDK configuration.
  */
  static Configuration  = class {
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    static UsageError = class {
      /**
       * This message indicates that a configuration option was not of the correct type. This primarily applies to languages that are not strongly typed.
      */
      static WrongType = class {
        static readonly code = "0:3:1";
        /**
         * Generate a log string for this code.
         * 
         * This function will automatically include the log code.
         * @param name The name of the configuration option.
         * @param expectedType The correct type for the configuration option.
         * @param actualType The incorrect types used for the configuration option.
         * @param defaultValue The default value of the configuration option.
        */
        static message(name: string, expectedType: string, actualType: string, defaultValue: string): string {
          return `0:3:1 Config option "${name}" should be of type ${expectedType}, but received ${actualType}, using the default value (${defaultValue}).`;
        }
      }
      /**
       * This message indicates that an unrecognized configuration option was provided. This primarily applied to languages that are not strongly typed.
      */
      static UnknownOption = class {
        static readonly code = "0:3:2";
        /**
         * Generate a log string for this code.
         * 
         * This function will automatically include the log code.
         * @param name The option that was not recognized.
        */
        static message(name: string): string {
          return `0:3:2 Ignoring unknown config option "${name}"`;
        }
      }
    }
  }
}
