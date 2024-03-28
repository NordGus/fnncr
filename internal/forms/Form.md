# Form

A form is the primary way how the system communicate in detail with the
outside, it represents the values that compose an
[Entity](../entities/Entity.md).

## Uses
- Insert new [Entities](../entities/Entity.md) in the system.
- Update an existing [Entity](../entities/Entity.md) in the system.
- Read the details of a persisted [Entity](../entities/Entity.md) in the
  system.

## Development Practices
- Use [Forms](#Form) as [DTOs](../../docs/concepts/Data%20Transfer%20Object.md)
  to communicate [Entity](../entities/Entity.md) state between the outside
  world and the system.
- When you are using a [Form](#Form) inside your system commands always
  validate that has been initialized correctly using its method
  `[form].Initialized()` it only returns `true` if the form was initialized
  correctly.
- If you need to add a new [Form](#Form) to the system please run:
    ```shell
    make generate form [yourformname] value1:type value2:type:nullable
    ```
- If you need to modify the generated code follow the following criteria:
  - Do not remove the `[form].Initialized()` method.
  - Do not remove nor modify the `initialized` field.
  - When you are initializing a form to insert or update an entity in the
    system, it should be initialized using a function with the pattern
    `New[YourImplementation]Entry`. And it requires a [raw form](#raw-form)
    as its only parameter. And make sure to set the `initialized` field to
    `true`.
  - When you are initializing a form to read an entity in the system, it should
    be initialized using a function with the pattern `New[YourImplementation]`.
    It requires an [Entity interface](#entity-interface) as its only parameter.
    And make sure to set the `initialized` field to `true`.
  - [Forms](#form) initialized from an [Entity interface](#entity-interface)
    is assumed to be a perfectly valid Form, so it doesn't need any validation
    whatsoever.
  - Make sure to write your [validation functions](#validation-function) for
    all form fields you need to validate, inside the `validations.go` file your
    form package directory.

## Glossary
### Raw Form
Not initialized using an initialization method in your form package. This would
usually happen when parsing the JSON request body inside the http rest adapter
or any other part you are building data to introduce to your system.

> Don't use a raw form inside your systems, always convert them into properly
> initialized forms,

### Entity Interface
Simple interface that represents the [Entity](../entities/Entity.md) the forms
is supposed to be a [DTO](../../docs/concepts/Data%20Transfer%20Object.md) of.

This is design decision is done to prevent circular dependencies. And also
facilitate refactoring.

### Validation Function
Function that should make a single validation on a `form_value` inner value and
return an error if the value doesn't pass it. Each function is supposed to run a
just on validation on the value, no more.

# TODO
- [x] Write Glossary definitions.
- [x] Refactor current implementation to match this spec.
- [X] Implement the form generation command/script.
