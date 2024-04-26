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
        static readonly code = "0:3:0";
        /**
         * Generate a log string for this code.
         * 
         * This function will automatically include the log code.
         * @param actualType The incorrect types used for the configuration option.
         * @param defaultValue The default value of the configuration option.
         * @param expectedType The correct type for the configuration option.
         * @param name The name of the configuration option.
        */
        static message(actualType: string, defaultValue: string, expectedType: string, name: string): string {
          return `0:3:0 Config option "${name}" should be of type ${expectedType}, but received ${actualType}, using the default value (${defaultValue}).`;
        }
      }
      /**
       * This message indicates that an unrecognized configuration option was provided. This primarily applied to languages that are not strongly typed.
      */
      static UnknownOption = class {
        static readonly code = "0:3:1";
        /**
         * Generate a log string for this code.
         * 
         * This function will automatically include the log code.
         * @param name The option that was not recognized.
        */
        static message(name: string): string {
          return `0:3:1 Ignoring unknown config option "${name}"`;
        }
      }
    }
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior. Setting the timeout for identification too high.
    */
    static UsageWarning = class {
    }
  }
  /**
   * Conditions related to polling.
  */
  static Polling  = class {
    /**
     * An error which represents a mis-use of an API and impedes correct functionality.
    */
    static UsageError = class {
      /**
       * You did a bad polling thing.
      */
      static BadPollThing = class {
        static readonly code = "1:3:0";
        /**
         * Generate a log string for this code.
         * 
         * This function will automatically include the log code.
         * @param bad The bad thing.
         * @param blank The blank thing.
        */
        static message(bad: string, blank: string): string {
          return `1:3:0 The ${blank} was how ${bad}.`;
        }
      }
    }
    /**
     * A warning about the usage of an API or configuration. The usage or configuration does not interfere with operation, but is not recommended or may result in unexpected behavior. Setting the timeout for identification too high.
    */
    static UsageWarning = class {
    }
  }
}
