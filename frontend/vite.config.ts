import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [svelte()],
    //build: {
    //    rollupOptions: {
    //        input: {
    //            main: resolve(__dirname, "/index.html"),
    //            //devices: resolve(__dirname, "/devices/index.html"),
    //        },
    //    },
    //},
    server: {
        proxy: {
            "/api": {
                target: "http://localhost:50833",
                changeOrigin: true,
            },
        },
    },
});
