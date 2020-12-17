with open("input.txt") as file:
    data = file.read().splitlines()

import numpy as np

def rotate(orientation, degrees):
    theta = np.radians(degrees)
    c, s = np.cos(theta), np.sin(theta)
    R = np.array(((c, -s), (s,c)))
    new_orient = np.matmul(R, orientation)
    return new_orient

# Part 1
starting_orient = np.array([[1], [0]])
starting_pos = np.array([[0], [0]])
pos = starting_pos.copy()
orient = starting_orient.copy()
for line in data:
    if line[0] == "N":
        pos[1] = pos[1] + int(line[1:])
    elif line[0] == "S":
                pos[1] = pos[1] - int(line[1:])
    elif line[0] == "E":
        pos[0] = pos[0] + int(line[1:])
    elif line[0] == "W":
        pos[0] = pos[0] - int(line[1:])
    elif line[0] == "L":
        orient = rotate(orient, +int(line[1:]))
    elif line[0] == "R":
        orient = rotate(orient, -int(line[1:]))
    elif line[0] == "F":
        pos = pos + int(line[1:]) * orient

man_dist = abs(pos[0]) + abs(pos[1])
print("manhattan distance = ", man_dist[0])

# Part 2
starting_orient = np.array([[1], [0]])
starting_pos = np.array([[0], [0]])
pos = starting_pos.copy()
orient = starting_orient.copy()
starting_waypoint = np.array([[10], [1]])
way_pos = starting_waypoint.copy()

for line in data:
    if line[0] == "N":
        way_pos[1] = way_pos[1] + int(line[1:])
    elif line[0] == "S":
                way_pos[1] = way_pos[1] - int(line[1:])
    elif line[0] == "E":
        way_pos[0] = way_pos[0] + int(line[1:])
    elif line[0] == "W":
        way_pos[0] = way_pos[0] - int(line[1:])
    elif line[0] == "L":
        way_pos = rotate(way_pos, +int(line[1:]))
    elif line[0] == "R":
        way_pos = rotate(way_pos, -int(line[1:]))
    elif line[0] == "F":
        pos = pos + int(line[1:]) * way_pos

man_dist = abs(pos[0]) + abs(pos[1])
print("manhattan distance = ", man_dist[0])
