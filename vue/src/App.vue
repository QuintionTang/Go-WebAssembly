<template>
    <div id="app">
        <Header />
        <div class="page-content">
            <div class="container">
                <div class="row">
                    <div class="col-lg-12">
                        <div class="text-center mt-sm-5 mb-4 text-white-80">
                            <div class="page-title">
                                使用 VUE 和 Go 触摸 WebAssembly
                            </div>
                            <p class="mt-3 fs-15">
                                WebAssembly（wasm）是指一种可以在浏览器及其外围技术和工具中运行的编程语言。它以二进制格式表示并由堆栈机器实现处理。
                            </p>
                        </div>
                    </div>
                </div>
                <div class="row mt-6">
                    <div class="col-lg-12">
                        <div class="card">
                            <div class="card-body">
                                <ul class="nav-pill">
                                    <li
                                        class="nav-item"
                                        :class="{
                                            actived: item.name === currentEg,
                                        }"
                                        v-for="item in arrayNavs"
                                        :key="item.name"
                                    >
                                        <a
                                            href="javascript:;"
                                            @click="changeEg(item)"
                                        >
                                            {{ item.title }}
                                        </a>
                                    </li>
                                </ul>
                                <div class="row">
                                    <div
                                        class="col-xl-4"
                                        v-if="currentEg === 'helloworld'"
                                    >
                                        <Helloworld />
                                    </div>
                                    <div
                                        class="col-xl-4"
                                        v-if="currentEg === 'helloresult'"
                                    >
                                        <HelloResult />
                                    </div>

                                    <div
                                        class="col-xl-4"
                                        v-if="currentEg === 'eventsresult'"
                                    >
                                        <EventsResult />
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <Footer />
    </div>
</template>

<script>
import Helloworld from "./components/Helloworld.vue";
import HelloResult from "./components/HelloResult.vue";
import EventsResult from "./components/Events.vue";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
export default {
    name: "App",
    components: {
        Helloworld,
        Header,
        HelloResult,
        Footer,
        EventsResult,
    },
    data() {
        const arrayNavs = [
            {
                title: "HelloWorld",
                name: "helloworld",
            },
            {
                title: "HelloResult",
                name: "helloresult",
            },
            {
                title: "Click Callback",
                name: "eventsresult",
            },
        ];
        return {
            arrayNavs,
            fpsInfo: null,
            metadata: null,
            currentEg: "helloworld",
        };
    },
    mounted() {
        this.init();
    },
    methods: {
        init() {},
        changeEg({ name }) {
            this.currentEg = name;
        },
    },
};
</script>
