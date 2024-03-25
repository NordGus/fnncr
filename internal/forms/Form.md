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
to communicate [Entity](../entities/Entity.md) between the outside world and
the system.
- When you are using a [Form](#Form) inside your system commands always 
validate that has been initialized correctly using its method
`[form].Initialized()` it only returns `true` if the form was initialized
correctly.
- If you need to add a new [Form](#Form) to the system please run:
    ```shell
    make generate form [yourformname] value1:type value2:type:nullable
    ```
- If you need to modify the generated code follow the following criteria:
  - Do not remove nor modify the `[form].Initialized()` method.
  - Do not remove nor modify the `initialized` field.
  - When you are initializing a form to insert or update an entity in the
  system, it should be initialized using a function with the pattern
  `New[YourImplementation]Entry`. And it requires a [raw form]()
  and a [Validator]() for the implementation as its parameters. And make sure
  to set the `initialized` field to `true`.
  - For a [Validator]() implementation use a struct. If the validator doesn't
  require any special setup keep the struct empty.
  - Write a validator create a file called `[yourimplementation]_validator.go`.
  - When you are initializing a form to read an entity in the system, it should
  be initialized using a function with the pattern `New[YourImplementation]`.
  It requires an [Entity interface]() as its only parameter. And make sure to
  set the `initialized` field to `true`.