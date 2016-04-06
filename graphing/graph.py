import json
from pprint import pprint
from ete3 import Tree, TreeStyle, TextFace


def parseTree(root):
    tree = Tree()
    tree.name = root['Name']
    tree.add_face(TextFace(root['Split']), column=0, position="branch-bottom")
    if root['Children']:
        for child in root['Children']:
            tree.children.append(parseTree(child))
    return tree


'''
with open('sample_tree.json', 'w') as outfile:
    json.dump(obj, outfile, sort_keys=True, indent=4, separators=(',', ': '))
'''

ts = TreeStyle()
ts.show_leaf_name = False

root = json.loads(open('test_tree.json').read())

pprint(root)
tree_root = parseTree(root)
print tree_root

for child in tree_root.traverse():
    # add a marker with the name of each node, at each node
    child.add_face(TextFace(child.name), column=0, position="branch-top")

# render the file and save it
tree_root.render("test.png", tree_style=ts, w=5000)
