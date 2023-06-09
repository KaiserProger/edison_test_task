<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <RouterLink
        to="/"
        custom
        v-slot="{ href, navigate }"
      >
        <a class="navbar-brand" :href="href" @click="navigate">Edison Extrasense</a>
      </RouterLink>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <RouterLink
            to="/history"
            custom
            v-slot="{ href, navigate }"
          >
            <li class="nav-item">
              <a class="nav-link active" aria-current="page" :href="href" @click="navigate">Guesses history</a>
            </li>
          </RouterLink>
          <RouterLink
            to="/trust"
            custom
            v-slot="{ href, navigate }"
          >
            <li class="nav-item">
              <a class="nav-link" :href="href" @click="navigate">Extrasenses trust</a>
            </li>
          </RouterLink>
        </ul>
      </div>
    </div>
  </nav>
</template>