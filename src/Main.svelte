<script lang="ts">
  import { Router, Route } from "@wjfe/n-savant";
  import type { SvelteComponentTyped } from "svelte";

  const modules = import.meta.glob("./routes/**/page.svelte");

  type Loader = () => Promise<{ default: typeof SvelteComponentTyped }>;

  const routes = Object.entries(modules).map(([path, loader]) => {
    let routePath = path.replace("./routes", "").replace("/page.svelte", "");

    if (routePath === "") routePath = "/";

    return { path: routePath, loader: loader as Loader };
  });
</script>

<Router>
  {#each routes as { path, loader }}
    <Route {path} key={path}>
      {#await loader() then mod}
        <svelte:component this={mod.default} />
      {/await}
    </Route>
  {/each}
</Router>
