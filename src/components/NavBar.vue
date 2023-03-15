<template>
  <div class="mx-auto max-w-7xl navbar bg-base-100">
    <div class="navbar-start">
      <RouterLink to="/" class="text-xl hover:text-accent">Noteballs</RouterLink>
    </div>
    <div class="hidden text-lg md:inline-flex navbar-end">
      <RouterLink to="/" class="rounded-md btn btn-ghost hover:text-accent" active-class="text-accent">Home</RouterLink>
      <div class="divider divider-vertical">|</div>
      <RouterLink to="/stats" class="btn btn-ghost hover:text-accent" active-class="text-accent">Stats</RouterLink>
    </div>
    <div class="md:hidden navbar-end">
      <div class="dropdown dropdown-end">
        <label tabindex="0" class="btn btn-ghost btn-circle">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" />
          </svg>
        </label>
        <ul tabindex="0" class="p-2 mt-3 w-40 rounded-xl shadow-xl menu menu-compact dropdown-content bg-base-100">
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
