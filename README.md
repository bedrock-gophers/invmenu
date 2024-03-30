# Inv
The Inv library allows the use of inventory menus, providing tools to create or send fake inventories.

## Creating an Inventory Menu
To create an inventory menu using the Inv library, follow these steps:

Place Fake Chest: Use the inv.PlaceFakeContainer function to place a fake chest, essential for the library to function properly.
```go
// Only run this once, after server start.
inv.PlaceFakeContainer(myWorld, cube.Pos{x, y, z})
```

Handle Player Packets: inv requires you to use the RedirectPlayerPackets on player join whichs makes it possible
for the library to handle incoming player packets, mostly use to handle container closing.
```go
// The 'p' variable represents the targeted player.
inv.RedirectPlayerPackets(p)
```

Create Menu Submittable: Your menu requires a menu submittable, here's an example:
```go
type MySubmittable struct{}

func (m MySubmittable) Submit(p *player.Player, it item.Stack) {
	fmt.Println("Submitted", it)
}

func (m MySubmittable) Close(p *player.Player) {
	fmt.Println("Closed")
}
```
Create Menu: Use the inv.NewMenu to create a new inventory menu.
```go
m := inv.NewMenu(MySubmittable{}, "test", inv.ContainerTypeChest)
```
Populate Menu Slots: Provide the menu with a `item.Stack` slice:
```go
var stacks = make([]item.Stack, 27)
for i := 0; i < 27; i++ {
    stacks[i] = item.NewStack(block.StainedGlass{Colour: item.ColourRed()}, 1)
}
m = m.WithStacks(stacks...)
```
Sending the Menu to a Player
To display the inventory menu to a player, use the inv.SendMenu function. Here's an example:

```go
// The 'p' variable represents the targeted player.
inv.SendMenu(p, m)
```
This code sends the inventory menu to the specified player.
