const updateActiveNav = () => {
  const sidebarItems = document.querySelectorAll("#sidebar-item")
  const navbarItems = document.querySelectorAll("#navbar-item")

  const current = window.location.pathname.replace(/\/+$/, "") || "/";

  // sidebar section
  const activeSibebarItem = ["bg-[#f5f5f5]", "dark:bg-[#262626]", "font-semibold"]
  sidebarItems.forEach(element => {
    element.classList.remove(...activeSibebarItem)
  })


  sidebarItems.forEach(element => {
    let navURL = element.getAttribute("href")
    if (!navURL) return

    navURL = navURL.replace(/[?#].*$/, "");
    console.log(navURL)

    if (current.includes(navURL)) {
      element.classList.add(...activeSibebarItem)
    }
  })

  // navbar section
  const activeNavbarItem = ["text-[#171717]", "dark:text-[#fafafa]"]
  const inactiveNavbarItem = ["text-neutral-500", "dark:text-neutral-500"]
  navbarItems.forEach(element => {
    element.classList.remove(...activeNavbarItem)
    element.classList.add(...inactiveNavbarItem)
  })

  navbarItems.forEach(element => {
    let navURL = element.getAttribute("href")
    if (!navURL) return

    navURL = navURL.replace(/\/+$/, "") || "/";


    if (navURL.includes(current)) {
      element.classList.remove(...inactiveNavbarItem)
      element.classList.add(...activeNavbarItem)
    }
  })
}

(() => {
  console.log("Welcome to Mystreams")
  document.addEventListener("DOMContentLoaded", updateActiveNav);
  document.body.addEventListener("htmx:pushedIntoHistory", updateActiveNav);
	window.addEventListener("popstate", updateActiveNav);
})();


