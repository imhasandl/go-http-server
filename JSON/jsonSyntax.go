// Assignment
// Complete the itemList and playerObject string constant values.

// itemList is a JSON array containing two items (object literals). It should have the following items:

// ITEM ONE:
//   id: 0 (number)
//   name: "sword" (string)
//   damage: 10.5 (number)
//   equipped: false (boolean)

// ITEM TWO:
//   id: 1 (number)
//   name: "shield" (string)
//   block: 5.5 (number)
//   equipped: true (boolean)
// Copy icon
// playerObject is a JSON object literal that represents a player. It should have the following fields:

// name: "Fudd" (string)
// items: "spear and magic helmet" (string)
// wife: "Brunhilde" (string)
// power: 9000 (number)
// Copy icon
// The tests simply check that the JSON is valid.


package main

const itemList = `
[
	{
		"id": 0,
		"name": "sword",
		"damage": 10.5,
		"equipped": false
	},
	{
		"id": 1,
		"name": "shield",
		"block": 5.5,
		"equipped": true
	}
]
`

const playerObject = `
{
	"name": "Fudd",
	"items": "spear and magic helmet",
	"wife": "Brunhilde",
	
	"power": 9000
}
`
