# Sugar Backend

 ## Table of Contents
- [Getting Started](#getting-started)
    - [Running with Docker](#running-with-docker)
    - [API Endpoints](#api-endpoints)
      - [User Endpoints](#user-endpoints)
      - [Product Endpoints](#product-endpoints)
      - [Note Endpoints](#note-endpoints)
    - [Schemas](#schemas)
      - [User Schemas](#user-schemas)
      - [Product Schemas](#product-schemas)
      - [Note Schemas](#note-schemas)

    ## Getting Started

    To start using the Sugar Backend API, you need to set up the environment and run the server. Follow these steps:

    1. **Clone the repository**:
       ```bash
       git clone https://github.com/lilpipidron/sugar-backend
       cd sugar-backend
       ```

    2. **Install dependencies**:
       ```bash
       go mod download
       ```

    3. **Run the server**:
       ```bash
       go run ./cmd/sugar-backend/main.go
       ```

    Your API server should now be running at `http://localhost:8080`.

    ## Running with Docker

    To run the Sugar Backend using Docker, you need Docker and Docker Compose installed on your system. Follow these steps:

    1. **Clone the repository**:
       ```bash
       git clone https://github.com/lilpipidron/sugar-backend.git
       cd sugar-backend
       ```

    2. **Build and start the containers**:
       ```bash
       docker-compose up --build
       ```

    This will start the Sugar Backend server on `http://localhost:8080` and a PostgreSQL database on `localhost:5432`.

    ### Docker Configuration

    #### Dockerfile
    ```dockerfile
    FROM golang:latest

    WORKDIR /app

    COPY . /app

    RUN go mod download

    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sugar-backend ./cmd/sugar-backend/main.go

    EXPOSE 8080

    CMD ["./sugar-backend"]
    ```

    #### docker-compose.yaml
    ```yaml
    version: "3.9"

    services:
      sugar-backend:
        build: ./
        command: ./sugar-backend
        ports:
          - "8080:8080"
        depends_on:
          - sugar-db
        environment:
          - POSTGRES_PASSWORD=postgres

      sugar-db:
        restart: unless-stopped
        image: postgres:latest
        ports:
          - "5432:5432"
        environment:
          - POSTGRES_PASSWORD=postgres
          - POSTGRES_USER=postgres
          - POSTGRES_DB=sugar-db
        volumes:
          - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    ```

    ## API Endpoints

    ### User Endpoints

    #### Get User
    - **GET** `/user`
    - **Description**: Fetch a user by login and password.
    - **Parameters**:
      - `login` (string) - User's login
      - `password` (string) - User's password
    - **Responses**:
      - `200 OK`

    #### Create User
    - **POST** `/user`
    - **Description**: Create a new user.
    - **Request Body**:
      - `AddUser` schema
    - **Responses**:
      - `201 Created`

    #### Delete User
    - **DELETE** `/user`
    - **Description**: Delete a user by ID.
    - **Request Body**:
      - `DeleteUser` schema
    - **Responses**:
      - `200 OK`

    #### Update User Information
    - **PUT** `/user/birthday`
      - **Description**: Change user's birthday.
      - **Request Body**:
        - `ChangeBirthday` schema
      - **Responses**:
        - `200 OK`

    - **PUT** `/user/breadUnit`
      - **Description**: Change user's bread unit.
      - **Request Body**:
        - `ChangeBreadUnit` schema
      - **Responses**:
        - `200 OK`

    - **PUT** `/user/carbohydrateRatio`
      - **Description**: Change user's carbohydrate ratio.
      - **Request Body**:
        - `ChangeCarbohydrateRatio` schema
      - **Responses**:
        - `200 OK`

    - **PUT** `/user/gender`
      - **Description**: Change user's gender.
      - **Request Body**:
        - `ChangeGender` schema
      - **Responses**:
        - `200 OK`

    - **PUT** `/user/name`
      - **Description**: Change user's name.
      - **Request Body**:
        - `ChangeName` schema
      - **Responses**:
        - `200 OK`

    - **PUT** `/user/weight`
      - **Description**: Change user's weight.
      - **Request Body**:
        - `ChangeWeight` schema
      - **Responses**:
        - `200 OK`

    ### Product Endpoints

    #### Get Products
    - **GET** `/product`
    - **Description**: Fetch products by name.
    - **Parameters**:
      - `name` (string) - Product name
    - **Responses**:
      - `200 OK`

    #### Get Carbs Amount
    - **GET** `/product/carbs`
    - **Description**: Fetch carbohydrate amount for a product by name.
    - **Parameters**:
      - `name` (string) - Product name
    - **Responses**:
      - `200 OK`

    #### Add Product
    - **POST** `/product`
    - **Description**: Add a new product.
    - **Request Body**:
      - `AddProduct` schema
    - **Responses**:
      - `201 Created`

    ### Note Endpoints

    #### Get Notes
    - **GET** `/note`
    - **Description**: Fetch notes by user ID.
    - **Parameters**:
      - `id` (integer) - User ID
    - **Responses**:
      - `200 OK`

    #### Get Notes by Date
    - **GET** `/note/date`
    - **Description**: Fetch notes by user ID and date.
    - **Parameters**:
      - `id` (integer) - User ID
      - `date-time` (string) - Date and time in ISO 8601 format
    - **Responses**:
      - `200 OK`

    #### Add Note
    - **POST** `/note`
    - **Description**: Add a new note.
    - **Request Body**:
      - `AddNote` schema
    - **Responses**:
      - `201 Created`

    #### Delete Note
    - **DELETE** `/note`
    - **Description**: Delete a note by note ID.
    - **Request Body**:
      - `DeleteNote` schema
    - **Responses**:
      - `200 OK`

    ## Schemas

    ### User Schemas

    #### AddUser
    ```yaml
    type: object
    properties:
      login:
        type: string
      password:
        type: string
      name:
        type: string
      birthday:
        type: string
        format: date-time
      gender:
        type: string
      weight:
        type: integer
      carbohydrate-ratio:
        type: integer
      bread-unit:
        type: integer
    ```

    #### ChangeBirthday
    ```yaml
    type: object
    properties:
      id:
        type: integer
      new_birthday:
        type: string
        format: date-time
    ```

    #### ChangeBreadUnit
    ```yaml
    type: object
    properties:
      id:
        type: integer
      new_bread_unit:
        type: integer
    ```

    #### ChangeCarbohydrateRatio
    ```yaml
    type: object
    properties:
      id:
        type: integer
      new_carbohydrate_ratio:
        type: integer
    ```

    #### ChangeGender
    ```yaml
    type: object
    properties:
      id:
        type: integer
      new_gender:
        type: string
    ```

    #### ChangeName
    ```yaml
    type: object
    properties:
      id:
        type: integer
      new_name:
        type: string
    ```

    #### ChangeWeight
    ```yaml
    type: object
    properties:
      id:
        type: integer
      new_weight:
        type: integer
    ```

    #### DeleteUser
    ```yaml
    type: object
    properties:
      id:
        type: integer
    ```

    ### Product Schemas

    #### AddProduct
    ```yaml
    type: object
    properties:
      id:
        type: integer
      name:
        type: string
      carbs:
        type: integer
    ```

    ### Note Schemas

    #### AddNote
    ```yaml
    type: object


    properties:
      user-id:
        type: integer
      note-id:
        type: integer
      note-type:
        type: string
      date-time:
        type: string
        format: date-time
      sugar-level:
        type: integer
      products:
        type: array
        items:
          type: object
          properties:
            id:
              type: integer
            name:
              type: string
            carbs:
              type: integer
    ```

    #### DeleteNote
    ```yaml
    type: object
    properties:
      id:
        type: integer
    ```