<template>
    <div class="card">
        <div class="card-header">
            <h4>Events</h4>
        </div>
        <div class="card-body">
            <p class="text-muted">
                点击“运行”，响应 Go WebAssembly 中绑定的事件，更改按钮文本。
            </p>
            <div class="live-preview">
                <button class="btn btn-success" id="clickresult">运行</button>
            </div>
        </div>
    </div>
</template>

<script>
import { initGo } from "./helper";
export default {
    name: "EventsResult",
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
                fetch("/wasms/events/main.wasm"),
                go.importObject
            )
                .then((result) => {
                    this.mod = result.module;
                    this.inst = result.instance;
                    this.run();
                })
                .catch((err) => {
                    console.error(err);
                });
        },
        async run() {
            await this.go.run(this.inst);
            this.inst = await WebAssembly.instantiate(
                this.mod,
                this.go.importObject
            );
        },
    },
};
</script>
