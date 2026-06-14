function toggleNav() {
	const nav = getElement("nav");
	const main = getElement("main");
	if (!nav || !main) {
		return;
	}

	
	if (nav.className === "hidden") {
		nav.className = "nav";
		main.className = "hidden";
	} else {
		nav.className = "hidden";
		main.className = "main";
	}
}

function getElement(id) {
	let e = document.getElementById(id);
	if (!e) {
		e = document.getElementsByClassName(id)[0];
		if (!e) {
			e = document.getElementsByTagName(id)[0];
			if (!e) {
				return null;
			}
		}
	}

	return e;
}
