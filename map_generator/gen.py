import random
import math
import noise
import numpy as np
from PIL import Image, ImageDraw
from matplotlib.colors import ListedColormap
from collections import deque
import sys


map_x, map_y = int(sys.argv[1]), int(sys.argv[1])
player_count = int(sys.argv[3])
harbor_count = player_count * 2
base_range = 5


def generate_perlin_noise(width, height, scale, octaves, persistence, lacunarity, seed):
    """Perlin noise generator"""
    world_map = np.zeros((width, height))
    island_centers = []

    for i in range(width):
        for j in range(height):
            x = i / scale
            y = j / scale
            value = noise.pnoise2(
                x,
                y,
                octaves=octaves,
                persistence=persistence,
                lacunarity=lacunarity,
                repeatx=1024,
                repeaty=1024,
                base=seed,
            )

            # Increase the threshold to generate more landmass
            if value > 0.2:
                world_map[i][j] = value
                island_centers.append((i, j))

    return world_map, island_centers


def find_harbor_coordinates(world_map):
    """Finds land tiles that have neighboring water tiles"""
    harbor_coordinates = set()
    coords = [(0, 1), (1, 0), (-1, 0), (0, -1)]
    visited = set()
    i, j = 0, 0
    start_coordinate = (i, j)
    queue = deque([start_coordinate])
    while queue:
        current_coordinate = queue.popleft()
        i, j = current_coordinate
        if current_coordinate not in visited:
            visited.add(current_coordinate)
            if is_land(world_map, i, j):
                for x, y in coords:
                    x, y = x + i, y + j
                    if 0 <= x < map_x and 0 <= y < map_y:
                        if is_water(world_map, x, y):
                            harbor_coordinates.add(current_coordinate)

            for x, y in coords:
                x, y = x + i, y + j
                if 0 <= x < map_x and 0 <= y < map_y and (x, y) not in visited:
                    queue.append((x, y))
    return harbor_coordinates


def fill_pixels(image, possible_coordinates, count, color, distance):
    """Fills the pixels on final image. Used for bases and harbors."""
    chosen_pixels = set()
    possible_coordinates = list(possible_coordinates)
    while len(chosen_pixels) < count:
        base_coordinate = random.choice(possible_coordinates)
        dist_ok = True
        for i in chosen_pixels:
            dist = math.hypot(base_coordinate[0] - i[0], base_coordinate[1] - i[1])
            if dist < distance:
                dist_ok = False
                break
        if dist_ok:
            chosen_pixels.add(base_coordinate)
            possible_coordinates.remove(base_coordinate)
            mark_color = color  # Red
            draw = ImageDraw.Draw(image)
            i, j = base_coordinate
            draw.point((i, j, i, j), fill=mark_color)
    return len(chosen_pixels)


def apply_colormap(data, colormap):
    normed_data = (data - data.min()) / (
        data.max() - data.min()
    )  # Normalize data to [0, 1]
    colored_data = (colormap(normed_data) * 255).astype(np.uint8)
    return colored_data


def is_water(colored_map, x, y):
    """Check if the pixel color represents water (e.g., blue)"""
    pixel_color = colored_map[y][x]
    return pixel_color[2] >= 128  # Assuming blue is above 128 in RGB


def is_land(colored_map, x, y):
    """Check if the pixel color represents land (e.g., green)"""
    pixel_color = colored_map[y][x]
    return pixel_color[1] >= 128  # Assuming green is above 128 in RGB


width, height = map_x, map_y  # Adjust the dimensions as needed
scale = 30.0  # Adjust the scale for Perlin noise (smaller scale for more islands)
octaves = 2
persistence = 0.4
lacunarity = 2.5
seed = np.random.randint(0, 250)

world_map, island_centers = generate_perlin_noise(
    width, height, scale, octaves, persistence, lacunarity, seed
)

# Create a custom colormap for land and ocean
custom_cmap = ListedColormap(["blue", "green"])

# Apply the colormap to the world map
colored_map = apply_colormap(world_map, custom_cmap)

# Convert the world map to an imagex
image = Image.fromarray(colored_map, mode="RGBA")

# Find and mark harbor coordinates as red dots
possible_coordinates = find_harbor_coordinates(colored_map)

print(
    "Generated",
    fill_pixels(image, possible_coordinates, harbor_count, (255, 0, 0, 255), 3),
    "harbor pixels",
)
print(
    "Generated",
    fill_pixels(
        image, possible_coordinates, player_count, (255, 255, 255, 255), base_range
    ),
    "base pixels",
)

# Display the image
# image.show()
image.save(sys.argv[4])
