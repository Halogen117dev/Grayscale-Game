components {
  id: "bulletScript"
  component: "/scripts/bulletScript.script"
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
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/bullets/bullet.atlas\"\n"
  "default_animation: \"bulletTravel\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -2.0
    y: -7.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    x: 0.499199
    y: 0.499199
    z: 4.991983
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullet\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -2.0\n"
  "      y: -8.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 7.3918695\n"
  "  data: 13.527564\n"
  "  data: 6.8\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
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
embedded_components {
  id: "sound1"
  type: "sound"
  data: "sound: \"/assets/music/gunshot/pistol_fire_handgun1.wav\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.155\n"
  "pan: 0.0\n"
  "speed: 1.2\n"
  "loopcount: 0\n"
  ""
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
embedded_components {
  id: "sound2"
  type: "sound"
  data: "sound: \"/assets/music/gunshot/pistol_fire_handgun2.wav\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.155\n"
  "pan: 0.0\n"
  "speed: 1.2\n"
  "loopcount: 0\n"
  ""
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
