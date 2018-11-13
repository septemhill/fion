# Fion

**Fion** is a log module for golang, provides different foreground color, background color, font control methods.

## Color, Font Control

**Fion** provides multiple color and font control methods to display output, as follows:

```golang
    //[Secion 1]
    fmt.Printf("%s\n", fion.Black("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.Red("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.Green("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.Yellow("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.Blue("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.Magenta("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.Cyan("%d %d", 100, 200))
    fmt.Printf("%s\n", fion.White("%d %d", 100, 200))

    //[Secion 2]
    fmt.Printf("%s %s\n", fion.Underline("Important"), fion.CrossedOut("Something Wrong"))

    //[Secion 3] Also you can use like this
    fmt.Printf("%s\n", fion.SlowBlink(fion.Red("Error!!!")))
```

In section 3, it's hard to use if you want use it on many place, so **Fion** offers you to custimize your own format:

```golang
    doublyUnderlineBlue := fion.NewFormat(fion.SDoublyUnderline, fion.SFDBlue)
    fmt.Println("%s\n", doublyUnderlineBlue("[WoW] Blue Doubly Underline"))
```

## Tag

There are few default tags, includes: 
- **Trace**, **Tracef**
- **Debug**, **Debugf**
- **Info**, **Infof**
- **Warn**, **Warnf**
- **Error**, **Errorf**

And **Fion** also provide customize method to create your own tag.

```golang
    //Use Fion provides tags
    fion.Trace("Where's Septem ?")
    fion.Debugf("Home with his %s.", "PS4")
    fion.Info("What's game he playing now ?")
    fion.Warnf("RDR%d", 2)
    fion.Errorf("loooooooooool")

    Notice, Noticef := fion.NewTag("[NOTICE]", 10, fion.SBBGreen, fion.SFBRed, fion.SInverse)
    Notice("Fion so Cooooooooooooooooooooool!!!")
    Noticef("u r running %x", 3735929054)
```

**Fion** tag also has level mechanism, default level setup to 2(**Info**), you can use SetTagLevel to change it.

```golang
    fion.Trace("Trace 0")       // Wouldn't show on your console
    fion.Debugf("Debug %d", 1)  // Wouldn't show on your console
    fion.Infof("Info %d", 2)    // Would show on your console
    fion.Warn("Warn 3")         // Would show on your console
    fion.Error("Error 4")       // Would show on your console

    //Let's change tag level, and run the same code section
    fion.SetTagLevel(0)

    fion.Trace("Trace 0")       // Would show on your console
    fion.Debugf("Debug %d", 1)  // Would show on your console
    fion.Infof("Info %d", 2)    // Would show on your console
    fion.Warn("Warn 3")         // Would show on your console
    fion.Error("Error 4")       // Would show on your console
```