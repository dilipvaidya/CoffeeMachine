# CoffeeMachine
Simple limited coffee machine implementation in go.

# Implementation Details
Most of the implementation changed at the later stage to support JSON parsing at the time of processing the same.

JSON string can be proovided as input same as sample provided 
# How to build (local build snapshot)
```
$ make build

$ ./bin/CoffeeMachine 
black_tea is prepared
hot_coffee is prepared
green_tea cannot be prepared because item sugar_syrup is 0
hot_tea cannot be prepared because item hot_water is 0

$ 
```
# Limitations
1. Detailing about which outlet the request is made for?
2. Monir changes for triggering the ingredient is below min level
3. Marking slots with ingredients to make have better design
4. performance - JSON parsing in evert request.
