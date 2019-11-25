<template>
    <div>
        <a-row>
            <div id="app">
                <div id="nav">
                    <router-link to="/">Home</router-link>
                    |
                    <router-link to="/Problem">Problem</router-link>
                    |
                    <router-link to="/Ace">Ace</router-link>
                    |
                    <router-link to="/About">About</router-link>
                </div>
            </div>
        </a-row>
        <a-row :gutter="8">
            <router-view/>
        </a-row>
    </div>
</template>

<style lang="scss">
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    #nav {
        padding: 30px;

    a {
        font-weight: bold;
        color: #2c3e50;

    &.router-link-exact-active {
        color: #42b983;
    }

    }
    }
</style>
