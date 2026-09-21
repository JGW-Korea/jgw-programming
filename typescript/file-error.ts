import fs from "fs";

function readFile(fileName: string) {
	const file = fs.readFileSync(fileName);


}

function exception() {
	readFile("undefined.txt");

	console.log("A");
}

function test() {
	try {
		exception();
	} catch(err) {
		if (err instanceof Error) {
			console.log(err.message)
		}
	}

	console.log("B");
}

test()