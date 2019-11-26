<template>
    <a-layout id="components-layout-demo-top" class="layout">
        <a-layout-header class="">
            <div class="logo" />
            <div id="app">
                <div id="nav" class="ant-menu ant-menu-horizontal ant-menu-root ant-menu-dark">
                    <router-link class="ant-menu-item" to="/">Home</router-link>
                    |
                    <router-link class="ant-menu-item" to="/Problems">Problems</router-link>
                    |
                    <router-link class="ant-menu-item" to="/Ace">Ace</router-link>
                    |
                    <router-link class="ant-menu-item" to="/About">About</router-link>
                </div>
            </div>
        </a-layout-header>
        <a-layout-content style="padding: 0 50px">
            <router-view :style="{ background: '#fff', padding: '24px', minHeight: '280px' }"/>
        </a-layout-content>
        <a-layout-footer style="text-align: center">
           Leaf ©2019 Created by Leaf
        </a-layout-footer>
    </a-layout>
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
        padding: 20px;

        a {
            font-weight: bold;
            color: #FFFFFF;

            &.router-link-exact-active {
                color: #42b983;
            }

        }
    }
</style>
