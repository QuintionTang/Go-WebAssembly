<template>
    <div class="card">
        <div class="card-header">
            <h4>HelloResult</h4>
        </div>
        <div class="card-body">
            <p class="text-muted">
                点击“运行”，将替换这里的内容
                <code id="helloresult">Hello, world!</code>
            </p>
            <div class="live-preview">
                <button
                    @click="run()"
                    :disbaled="loading"
                    class="btn btn-success"
                    id="runButton"
                    disabled
                >
                    <span v-if="loading"> 运行中…… </span>
                    <span v-else> 运行 </span>
                </button>
            </div>
        </div>
    </div>
</template>

<script>
import { initGo } from "./helper";
export default {
    name: "HelloResult",
    data() {
        return {
            go: null,
            mod: null,
            inst: null,
            result: "",
            loading: false,
        };
    },
    mounted() {
        this.init();
    },
    methods: {
        init() {
            if (!WebAssembly.instantiateStreaming) {
                // polyfill
                WebAssembly.instantiateStreaming = async (
                    resp,
                    importObject
                ) => {
                    const source = await (await resp).arrayBuffer();
                    return await WebAssembly.instantiate(source, importObject);
                };
            }

            const go = initGo();
            this.go = go;
            WebAssembly.instantiateStreaming(
                fetch("/wasms/document/main.wasm"),
                go.importObject
            )
                .then((result) => {
                    this.mod = result.module;
                    this.inst = result.instance;
                    document.getElementById("runButton").disabled = false;
                })
                .catch((err) => {
                    console.error(err);
                });
        },
        async run() {
            this.loading = true;
            await this.go.run(this.inst);
            this.inst = await WebAssembly.instantiate(
                this.mod,
                this.go.importObject
            );
            this.loading = false;
        },
    },
};
</script>
