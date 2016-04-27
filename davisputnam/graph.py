import json
import sys
from pprint import pprint
from ete3 import Tree, TreeStyle, TextFace
from PIL import Image, ImageDraw


def parseTree(root):
    tree = Tree()
    tree.name = root['Name']
    tree.add_face(TextFace(root['Split'], fgcolor="red"), column=0, position="branch-bottom")
    if root['Children']:
        for child in root['Children']:
            tree.children.append(parseTree(child))
    return tree


'''
with open('sample_tree.json', 'w') as outfile:
    json.dump(obj, outfile, sort_keys=True, indent=4, separators=(',', ': '))
'''

if __name__ == '__main__':
    ts = TreeStyle()
    ts.show_leaf_name = False

    root = json.loads(open(sys.argv[1]).read())

    pprint(root)
    tree_root = parseTree(root)
    print tree_root

    for child in tree_root.traverse():
        # add a marker with the name of each node, at each node
        child.add_face(TextFace(child.name), column=0, position="branch-top")

    # render the file and save it
    fname = sys.argv[1][:-4] + "png"
    tree_root.render(fname, tree_style=ts, w=5000)

    im = Image.open(fname)
    (x, y) = im.size

    draw = ImageDraw.Draw(im)
    draw.rectangle((0, y*.45, x*.25, y), fill="white")
    im.save(fname, "PNG")
    # tree_root.show(tree_style=ts)
