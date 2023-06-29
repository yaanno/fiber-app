window.addEventListener('load', function() {
  const input = document.getElementById("search");
  const grid = document.getElementById("table");
  const form = document.getElementById("form");
  form.addEventListener('submit', function(e) {
    e.preventDefault();
  })
  input.addEventListener('keyup', async function(e) {
    const fragment = await (await fetch(`/search?name=${e.target.value}`)).text()
    if (fragment.length > 0) {
      grid.innerHTML = fragment;
    }
  })
})
