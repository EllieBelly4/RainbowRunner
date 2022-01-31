const fs = require("fs")

const data = fs.readFileSync("../debug_sorted.txt")

let acc = 16

for (const line of data.toString().split("\n")) {
    if (line.trim() === "") {
        break;
    }

    const splitLine = line.split(" ")
    const offset = parseInt(splitLine[0])
    const len = parseInt(splitLine[1])

    if (acc !== offset) {
        console.log(`gap: ${offset} ${offset - acc} bytes`)
    }

    acc += len
}

console.log(acc)
