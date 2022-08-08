# Sway power button

A simple button for the swaybar implement written in Go with Wails

## WIP

### Custom config NOT IMPLEMENTED YET

#### This is a WIP what things could look like for now you must change the /frontend/src/style.css and rebuild the package yourself

Create a style.css file located at $HOME/.config/swaypowerbtn

Change the root's variables

Default config:

```css
/* style.css */
:root {
  --bgc: rgba(46, 52, 64, 1); // Change the background color
  --gap: 1rem; // Change the gap between buttons
  --border-radius: 2rem; // button radius
  --font-family: Arial, Helvetica, sans-serif; // font family used
  --txt-clr: rgba(216, 222, 233, 1); // svg/text color
  --btn-bgc: transparent; // button default background
  --btn-hover-bgc: rgba(
    216,
    222,
    233,
    0.6
  ); // button background-color on mouse hover
  --btn-hover-clr: rgba(76, 86, 106, 1) // button svg/text color on mouse hover ;
}
```

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

cmd:

```bash

make build

```
