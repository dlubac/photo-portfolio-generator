// Lightbox for gallery and archive pages: one configuration everywhere,
// captions from data-sub-html, focus returned to the photo you were on.
(function () {
	var sheet = document.getElementById("sheet");
	if (!sheet || typeof window.lightGallery !== "function") return;

	var links = sheet.querySelectorAll(".frame-link");
	var themeMeta = document.querySelector('meta[name="theme-color"]');
	var paper = themeMeta ? themeMeta.getAttribute("content") : null;
	var current = 0;

	var heading = document.querySelector("h1");
	var viewerName = (heading ? heading.textContent.trim() + ", " : "") + "photo viewer";

	sheet.addEventListener("lgBeforeOpen", function () {
		if (themeMeta) themeMeta.setAttribute("content", "#0d0e10");
	});

	sheet.addEventListener("lgAfterOpen", function () {
		var dialog = document.querySelector(".lg-container [role='dialog'], .lg-container");
		if (dialog) dialog.setAttribute("aria-label", viewerName);
	});

	sheet.addEventListener("lgAfterSlide", function (event) {
		current = event.detail.index;
	});

	sheet.addEventListener("lgAfterClose", function () {
		if (themeMeta && paper) themeMeta.setAttribute("content", paper);
		var link = links[current];
		if (link) {
			link.focus({ preventScroll: true });
			link.scrollIntoView({ block: "nearest" });
		}
	});

	window.lightGallery(sheet, {
		selector: ".frame-link",
		speed: 400,
		download: false,
		counter: true,
		loop: true,
		zoom: false,
		thumbnail: false,
		hideBarsDelay: 2500,
		// On touch screens, swipe moves between photos and swipe down closes;
		// the overlaid arrows are hidden so nothing sits on the image.
		mobileSettings: { controls: false, showCloseIcon: true, download: false },
	});
})();
