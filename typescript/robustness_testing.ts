function isNumber(value: unknown): value is number {
	if (typeof value === "number" && !Number.isNaN(value)) {
		return true;
	}

	return false;
}

/**
 * 두 개의 숫자를 전달받아 더한 결과를 반환하는 함수 (강건성 보장 O)
 *
 * @param   {number} x  연산에 사용할 첫 번째 피연산자
 * @param   {number} y  연산에 사용할 두 번째 피연산자
 * @returns {number}    두 피연산자를 더한 결과
 */
function add(x: unknown, y: unknown): number {
	
	
	if (!isNumber(x) || !isNumber(y)) {
		throw new TypeError(
			"Invalid argument: Both arguments must be valid numbers."
		);
	}
	
	if (!Number.isSafeInteger(x) || !Number.isSafeInteger(y)) {
		throw new RangeError("Argument exceeds safe integer range.");
	}
	
	let result: number = x + y;
	
	if (!Number.isSafeInteger(result)) {
		throw new RangeError("Result exceeds safe integer range.");
	}
	
	return result;
}