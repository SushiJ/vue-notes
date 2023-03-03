<template>
  <div class="navbar bg-base-100 max-w-7xl mx-auto">
    <div class="navbar-start">
      <RouterLink to="/" class="text-xl hover:text-accent">Noteballs</RouterLink>
    </div>
    <div class="navbar-end text-lg hidden md:inline-flex">
      <RouterLink to="/" class="btn btn-ghost hover:text-accent rounded-md" active-class="text-accent">Home</RouterLink>
      <div class="divider divider-vertical">|</div>
      <RouterLink to="/stats" class="btn btn-ghost hover:text-accent" active-class="text-accent">Stats</RouterLink>
    </div>
    <div class="navbar-end md:hidden">
      <div class="dropdown dropdown-end">
        <label tabindex="0" class="btn btn-ghost btn-circle">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" />
          </svg>
        </label>
        <ul tabindex="0" class="menu menu-compact dropdown-content mt-3 p-2 shadow-xl bg-base-100 rounded-xl w-40">
          <li>
            <RouterLink to="/" class="btn btn-accent btn-outline" active-class="text-accent">Home</RouterLink>
          </li>
          <li>
            <RouterLink to="/stats" class="btn btn-accent btn-outline" active-class="text-accent">Stats</RouterLink>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>
