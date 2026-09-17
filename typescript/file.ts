import fs from "fs";

function readFile(fileName: string) {
	const text = fs.readFileSync(fileName);

	if (text) {
		console.log(true);
		return;
	}

	console.log(false);
	return;
}

readFile("undefined.txt");