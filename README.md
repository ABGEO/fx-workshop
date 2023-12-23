# fx-workshop

This repository contains code from my GopherUp v1.0 meetup talk - `From Complexity to Clarity:
Crafting Seamless DI with Uber FX`.

The presentation could be found on [SlideShare](https://www.slideshare.net/TemuriTakalandze/from-complexity-to-clarity-crafting-seamless-di-with-uber-fx).

## Branches

- `main`: The initial, ugly and unrefactored version for demonstrating challenges of dependency management.
- `refactored`: The refactored version that utilizes Uber FX framework.

## Getting Started

- Clone this repository

    ```shell
    git clone https://github.com/ABGEO/fx-workshop.git
    ```

- After cloning, you can switch between the `main` and `refactored` branches to explore different versions of the
  project

    ```shell
    git checkout main      # For the initial, unrefactored version.
    git checkout refactored  # For the refactored version.
    ```

- Download Go modules

   ```shell
   go mod download
   ```

- Run the project

    ```shell
    go run .
    ```

## Authors

- [Temuri Takalandze](https://abgeo.dev)

## License

Copyright (c) 2023 [Temuri Takalandze](https://abgeo.dev).  
Released under the [GPL-3.0](LICENSE) license.
