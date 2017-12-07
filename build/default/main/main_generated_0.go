embedded_components {
  id: "label"
  type: "label"
  data: "size {\n  x: 128.0\n  y: 32.0\n  z: 0.0\n  w: 0.0\n}\nscale {\n  x: 2.0\n  y: 2.0\n  z: 1.0\n  w: 0.0\n}\ncolor {\n  x: 1.0\n  y: 1.0\n  z: 1.0\n  w: 1.0\n}\noutline {\n  x: 0.0\n  y: 0.0\n  z: 0.0\n  w: 1.0\n}\nshadow {\n  x: 0.0\n  y: 0.0\n  z: 0.0\n  w: 1.0\n}\nleading: 1.0\ntracking: 0.0\npivot: PIVOT_CENTER\nblend_mode: BLEND_MODE_ALPHA\nline_break: false\ntext: \"Greetings traveller!\"\nfont: \"/def-alien-world/fonts/text.font\"\nmaterial: \"/builtins/fonts/label.material\"\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/def-alien-world/hud/hud.atlas\"\ndefault_animation: \"textbox\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
