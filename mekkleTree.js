const crypto = require("crypto");

class MerkleTree {
  constructor(data) {
    this.data = data;
    this.leaves = this.data.map(this.hash);
    this.tree = this.createTree(this.leaves);
  }

  hash(data) {
    return crypto.createHash("sha256").update(data).digest("hex");
  }

  createTree(leaves) {
    const store = [leaves]; 
    const levels = Math.ceil(Math.log2(leaves.length));

    for (let level = 0; level < levels; level++) {
      const nextLevel = [];
      for (let i = 0; i < leaves.length; i += 2) {
        const left = leaves[i];
        const right = i + 1 < leaves.length ? leaves[i + 1] : left; 
        nextLevel.push(this.hash(left + right));
      }
      store.push(nextLevel);
      leaves = nextLevel;
    }

    return store;
  }

  getRoot() {
    return this.tree[this.tree.length - 1][0]; 
  }

  getProof(transactionToProof) {
    const hashTrx = this.hash(transactionToProof);
    let index = this.leaves.indexOf(hashTrx);
    if (index === -1) return [];

    const proof = [];
    for (let level = 0; level < this.tree.length - 1; level++) {
      const isRightNode = index % 2 === 1;
      const siblingIndex = isRightNode ? index - 1 : index + 1;

      if (siblingIndex < this.tree[level].length) {
        proof.push({
          direction: isRightNode ? "left" : "right",
          hash: this.tree[level][siblingIndex],
        });
      }

      index = Math.floor(index / 2); 
    }

    return proof;
  }

  verifyProof(root, transaction, proof) {
    let hash = this.hash(transaction);

    for (const { direction, hash: siblingHash } of proof) {
      if (direction === "left") {
        hash = this.hash(siblingHash + hash);
      } else if (direction === "right") {
        hash = this.hash(hash + siblingHash);
      }
    }

    return hash === root;
  }
}

const transactions = [
  "Victor",
  "Manji",
  "Ola",
  "Jerry",
  "Victory",
  "James",
  "Emeka",
  "Deborah",
  "Kenneth",
  "Richmond",
  "Ali",
];

const merkleTree = new MerkleTree(transactions);
const root = merkleTree.getRoot();
console.log("Merkle Root:", root);

const transactionToProof = "victor";
const proof = merkleTree.getProof(transactionToProof);
console.log(`Proof for "${transactionToProof}":`, proof);

const isValid = merkleTree.verifyProof(root, transactionToProof, proof);
console.log(`Is "${transactionToProof}" valid?`, isValid);
