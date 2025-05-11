<script module lang="ts">
    import type { Snippet } from "svelte";
    type PropsType = {
        children?: Snippet;
        class?: string;
        variant?: "normal" | "danger" | "help";
    };
</script>

<script lang="ts">
    import { Button } from "bits-ui";
    let {
        children,
        class: className,
        variant = "normal",
        ...rest
    }: PropsType = $props();

    const variantClasses = {
        normal: "bg-dark text-background hover:bg-dark/95",
        danger: "bg-red-600 text-white hover:bg-red-700",
        help: "bg-blue-500 text-white hover:bg-blue-600",
    };
</script>

<Button.Root
    class={`rounded-input shadow-mini inline-flex
    h-12 items-center justify-center px-[21px] text-[15px]
    font-semibold active:scale-[0.98] active:transition-all 
    ${variantClasses[variant]} ${className || ""}`}
    {...rest}
>
    {#if children}
        {@render children?.()}
    {/if}
</Button.Root>
