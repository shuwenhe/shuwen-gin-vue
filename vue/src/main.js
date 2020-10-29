import Vue from "vue"
import App from "./App.vue"
import router from "./router"
// import store from "./store"

Vue.config.productionTip = false

const name = "Richard"
    // name = "Ritchie";

function sayHello(who) {
    console.log(`hello${who}`)
}

sayHello(name)

new Vue({
    router,
    // aiaiiistore,
    render: h => h(App),
}).$mount("#app")
