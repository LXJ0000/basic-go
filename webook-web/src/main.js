import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

async function bootstrap() {
    const app = createApp(App)
    setupStore(app) // 优先级最高
    await setupRouter(app)
    app.mount('#app')
}
bootstrap()