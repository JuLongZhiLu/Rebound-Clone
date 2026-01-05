components {
  id: "building"
  component: "/scene3d/scripts/prefabs/building.script"
  properties {
    id: "frustum_mesh_max_dimension"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "collision"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.5\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "mask: \"player\"\n"
  "mask: \"ball\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -0.1\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 5.0\n"
  "  data: 0.1\n"
  "  data: 5.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "mesh"
  type: "mesh"
  data: "material: \"/scene3d/materials/basic_grid_timberwolf.material\"\n"
  "vertices: \"/scene3d/assets/meshes/buildings_floor_10x10.buffer\"\n"
  "textures: \"/scene3d/assets/textures/grid_20x20.png\"\n"
  "position_stream: \"position\"\n"
  "normal_stream: \"normal\"\n"
  ""
}
