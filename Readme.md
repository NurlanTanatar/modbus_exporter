## Modbus Data Pusher

This Go application retrieves data from a Modbus device, specifically a PLC (Programmable Logic Controller), and pushes that data to a Prometheus Pushgateway for monitoring and analysis.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [TODO](#todo)

## Installation

1. Clone the repository:

   ```shell
   git clone github.com/NurlanTanatar/modbus_exporter
   ```
2. Build the application:

   ```shell
   go build
   ```

## Usage

1. Ensure you have a `.env` file in the parent directory containing the necessary environment variables:

   ```env
   pushgw_host=<your_pushgateway_host>
   pushgw_user=<your_pushgateway_username>
   pushgw_pass=<your_pushgateway_password>
   ```
2. Run the application:

   ```shell
   ./modbus_data_pusher
   ```

## Configuration

- This application uses the [joho/godotenv](https://github.com/joho/godotenv) package to load environment variables from a `.env` file. Make sure to create this file in the parent directory of your application and populate it with the required variables as mentioned in the [Usage](#usage) section.
- The Modbus data is retrieved from a package named `data` in your project. Ensure that this package correctly fetches the Modbus data. The data is used to create Prometheus metrics.
- Prometheus metrics are created using the [prometheus/client_golang](https://github.com/prometheus/client_golang) package. Two metrics are created: `tic_100_temperature` and `current_time`. The former represents the temperature reading from the Modbus device, and the latter represents the timestamp when the data is collected.
- The metrics are then pushed to a Prometheus Pushgateway using the [prometheus/client_golang/prometheus/push](https://github.com/prometheus/client_golang/tree/master/prometheus/push) package. Ensure that you have a Prometheus Pushgateway set up and running with the correct credentials specified in your `.env` file.

## TODO

- [ ] Add error handling to gracefully handle failures during the Modbus data retrieval process.
- [ ] Implement more advanced data processing or filtering logic as needed for your specific use case.
- [ ] Create a Dockerfile for easy containerization and deployment of the application.
- [ ] Add unit tests to ensure the correctness and reliability of the code.
- [ ] Provide a more detailed explanation of the Modbus data retrieval process in the README, including any specific Modbus libraries or tools used.
- [ ] Consider adding additional documentation on how to set up and configure Prometheus and Prometheus Pushgateway for monitoring and analysis.
