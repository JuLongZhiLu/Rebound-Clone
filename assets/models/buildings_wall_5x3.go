components {
  id: "building"
  component: "/scene3d/scripts/prefabs/building.script"
  properties {
    id: "frustum_mesh_max_dimension"
    value: "5.0"
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
  "      y: 1.5\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 5.079381\n"
  "  data: 1.5\n"
  "  data: 0.05\n"
  "}\n"
  ""
}
