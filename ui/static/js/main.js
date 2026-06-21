function toggleNav() {
	const nav = getElement("nav");
	const main = getElement("main");
	const toggle = getElement("navToggle");
	if (!nav || !main || !toggle) {
		return;
	}

	
	if (nav.className === "hidden") {
		// show nav
		nav.className = "nav";
		main.className = "hidden";
		toggle.innerHTML = "-";
	} else {
		// hide nav
		nav.className = "hidden";
		main.className = "main";
		toggle.innerHTML = "X";
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
