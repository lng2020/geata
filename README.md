Geata is an IEC61850-based IoT gateway, which uses modern web tech stack including Vue3, TailwindCSS, Vite, Go, Gin and Xorm. It has serveral features like:
1. Converting MQTT and ModbusTCP data into IEC61850. Every data attribute(shorted as DA) in IEC61850 is configurable in web UI.
2. Fetching IEC61850 data from the server by using the self compiled client binary, which is generated from [libiec61850](https://github.com/mz-automation/libiec61850). The source code of the client binary is located in the `internal/app/client/client.c` file. I'm too lazy to build linux binary, you can build it by yourself.
3. Parsing ICD/SCD file and storing the data in the database.
4. Basic station and user management.

## Usage
1. copy the `config.example.yaml` to `config.yaml` and modify it.
2. frontend
```bash
cd web
pnpm install
pnpm run dev
```
3. backend
```bash
go build
./geata start
```
4. mock MQTT client
```bash
go build
./geata mock
```

## TODO and Limitation
- [ ] Add more configuration options for MQTT and ModbusTCP and adapt more data types.
- [ ] Currently `FC`(functional constraint in IEC61850) is hard coded in the IEC61850 client. It should be detected and recorded in DB.
- [ ] Performance optimization for the IEC61850 client and modbus client. It uses polling now and should be changed to event driven.
- [ ] Considering using IEC61850 go library to replace the self compiled client binary. But there is no go library for IEC61850 now...Without the library, the client binary is the only way to fetch data from the server and can't serve IE61850 data to the client.
- [ ] ICD/SCD file parsing is limited due to my lack of knowledge of IEC61850. It should be improved.

## Project structure
```
- cmd # main package and mock command
- internal
    - app # business logic
        - client # IEC61850 client
        - handler # protocol handler
        - parser # ICD/SCD file parser
        - model # data model
        - service # service layer
        - web # web and router
        - logger # logger
        config.go # configuration
        app.go # entry point
    - docs # documents auto generated by swaggo
- web # frontend
```

## How to adapt more protocols
Implement the interface `Handler` and `HandlerConfig` in the `internal/app/handler` directory and register it in the `internal/app/handler/handler.go` file.

## License
MIT