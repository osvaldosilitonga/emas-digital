# Jojonomic - Emas Digital

## How to Use

- Clone Repo:

  ```bash
  $ git clone https://github.com/osvaldosilitonga/emas-digital.git
  ```

- Run Docker Compose:

  ```bash
  $ docker-compose up
  ```

- Run Go Migrations:
  ```bash
  $ migrate -database "postgres://adminpostgres:87654321@localhost:5435/emasdigital?sslmode=disable" -path ./misc/migrations up
  ```

## Insomnia API Client

- Make sure you have Insomnia installed.

- Insomnia YAML file location

  ![Insomnia YAML File Location](./misc/docs/insomnia_file_location.png)

- Create new project

  ![New Project](./misc/docs/create_new_project.png)

- Upload insomnia YAML file

  ![Upload File](./misc/docs/choose_file.png)

- Import file

  ![Import File](./misc/docs/import.png)

  ![Scan File](./misc/docs/scan_file.png)

  ![After Import File](./misc/docs/after_import.png)

- Happy Testing

  ![Happy Testing](./misc/docs/happy_testing.png)
