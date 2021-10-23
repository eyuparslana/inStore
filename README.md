# In Store

A rest-api that works with golang as an in-memory key value store

## Usage

Fist of all, clone the repo with the command below. You must have golang installed on your computer to run the project.

```shell
git clone https://github.com/eyuparslana/inStore.git
```

+ ### Locally

Set the **`API_PORT`**, **`EXPORT_FILE_PATH`** and **`RECORD_FREQ`** environments.
And run command below in terminal in project directory.

```shell
go run main.go
```
Now, you can send requests to the application either via postman, swagger or browser (GET requests only) according to the following endpoint table.

+ ### With Docker

If you have docker installed on your system, you can run the project with docker.
Run command below on terminal.

````shell
docker-compose up -d
````

Now go to `localhost:8001` in your browser. You will see the swagger page. Here you can send your requests to api.

or you can test the api on the URL below.

**[InStoreApp](https://in-store-app.ew.r.appspot.com/)**.

## Endpoint Table

| Endpoints | Descriptions |  Methods |
| :------:| :-----------:| :-----------:|
| /   | returns all data contained in DB in key value format. | GET |
| /set | stores the key value in the JSON that comes with the POST request in the inMemDB key value format. |POST|
| /get |Get returns the value corresponding to 'key' in DB. | GET |
| /flush | Flush deletes all data in DB. | GET |


