<template>
    <div class="card">
        <div class="card-header">
            <h4>Helloworld</h4>
        </div>
        <div class="card-body">
            <p class="text-muted">
                点击“运行”，在控制台输出日志 <code>Hello, WebAssembly!</code>
            </p>
            <div class="live-preview">
                <button
                    @click="run()"
                    class="btn btn-success"
                    id="runButton"
                    disabled
                >
                    运行
                </button>
            </div>
        </div>
    </div>
</template>

<script>
import { initGo } from "./helper";
export default {
    name: "Helloworld",
    data() {
        return {
            go: null,
            mod: null,
            inst: null,
            result: "",
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
                fetch("/wasms/helloworld/main.wasm"),
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
            console.clear();
            await this.go.run(this.inst);
            this.inst = await WebAssembly.instantiate(
                this.mod,
                this.go.importObject
            );
        },
    },
};
</script>
