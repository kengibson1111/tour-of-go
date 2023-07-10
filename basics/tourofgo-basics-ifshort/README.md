# basics - if (short statement)

The `if` statement can have a "short assignment" init statement before the condition. Variable short assignment is covered in **basics - variables (short declarations)**. The variable being assigned is only in the `if` block's scope.

I'm not sure how effective this is as a coding practice. I can get the same effect moving the `if` init statement to one line above. The variable is in function scope, but I don't know if that is necessarily a negative thing in the runtime.

If I want to train people who are familiar with C or Java, it is more effective to keep the `if` statement as close as possible to representations in other languages.
