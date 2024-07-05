# Restaurant Order Management System

## Overview

This project is a Restaurant Order Management System implemented in Go. It allows users to view a menu, place orders, modify orders, and generate bills. The menu includes both Bengali and North Indian dishes.

## Project Structure

The project is divided into several Go files for better modularity and management:

- `main.go`: Entry point of the application.
- `menu.go`: Handles menu-related functionalities.
- `order.go`: Manages order-related functionalities.
- `utils.go`: Contains utility functions for displaying messages and processing orders.

## Methods and Their Descriptions

### Main Functions

- `main()`: The entry point of the application. It greets the user, shows the menu, takes orders, processes orders, and says goodbye.

### Menu Functions (`menu.go`)

- `getMenu()`: Initializes and returns the Bengali and North Indian menus.
- `showMenu(bengaliMenu, northIndianMenu *orderedmap.OrderedMap)`: Displays the Bengali and North Indian menus.

### Order Functions (`order.go`)

- `takeOrder(bengaliMenu, northIndianMenu *orderedmap.OrderedMap)`: Takes orders from the user.
- `modifyOrder(order *map[MenuItem]int)`: Allows the user to modify their order.
- `printBill(order map[MenuItem]int)`: Prints the bill for the current order.
- `updateOrder(update string, order *map[MenuItem]int)`: Updates the quantity of items in the order.
- `deleteItem(update string, order *map[MenuItem]int)`: Deletes an item from the order.

### Utility Functions (`utils.go`)

- `greet(customerName string)`: Greets the user.
- `sayBye(customerName string)`: Bids farewell to the user.
- `printDashedLine()`: Prints a dashed line for formatting the bill.
- `printGenerateBill()`: Prints the "Generating Bill" message.
- `orderProcessing()`: Simulates order processing.

## How to Use This Project

### Prerequisites

- Go installed on your machine.
- Make installed on your machine.

### Building and Running the Project

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/restaurant-oms.git
   cd restaurant-oms
   ```

2. Build the project:
   ```sh
   make build
   ```
3. Run the project:
   ```sh
   make run
   ```
4. Clean the build files:
   ```sh
   make clean
   ```

### Running the Project Manually

1. Build the project:

   ```sh
    go build -o restaurant main.go menu.go order.go utils.go
   ```

2. Run the project:
   ```sh
    ./restaurant
   ```

### Contribution

Contributions are welcome! If you have any suggestions, bug reports, or improvements, feel free to create an issue or submit a pull request.

### License

This project is licensed under the MIT License. See the [LICENSE file](https://opensource.org/license/mit) for details.

### Acknowledgements

Special thanks to the Go community for their continuous support and contributions.
