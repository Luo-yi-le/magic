<template>
    <div class="layout-columns-aside">
        <el-scrollbar>
            <ul>
                <li
                    v-for="(v, k) in columnsAsideList"
                    :key="k"
                    @click="onColumnsAsideMenuClick(v, k)"
                    :ref="
                        (el) => {
                            if (el) columnsAsideOffsetTopRefs[k] = el;
                        }
                    "
                    :class="{ 'layout-columns-active': liIndex === k }"
                    :title="v.meta.title"
                >
                    <div class="layout-columns-aside-li-box" v-if="!v.meta.link || (v.meta.link && v.meta.isIframe)">
                        <i :class="v.meta.icon"></i>
                        <div class="layout-columns-aside-li-box-title font12">
                            {{ v.meta.title && v.meta.title.length >= 4 ? v.meta.title.substr(0, 4) : v.meta.title }}
                        </div>
                    </div>
                    <div class="layout-columns-aside-li-box" v-else>
                        <a :href="v.meta.link" target="_blank">
                            <i :class="v.meta.icon"></i>
                            <div class="layout-columns-aside-li-box-title font12">
                                {{ v.meta.title && v.meta.title.length >= 4 ? v.meta.title.substr(0, 4) : v.meta.title }}
                            </div>
                        </a>
                    </div>
                </li>
                <div ref="columnsAsideActiveRef" :class="setColumnsAsideStyle"></div>
            </ul>
        </el-scrollbar>
    </div>
</template>

<script lang="ts">
import { reactive, toRefs, ref, computed, onMounted, nextTick, getCurrentInstance, watch } from 'vue';
import { useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router';
import { useStore } from '@/store/index.ts';
export default {
    name: 'layoutColumnsAside',
    setup() {
        const columnsAsideOffsetTopRefs: any = ref([]);
        const columnsAsideActiveRef = ref();
        const { proxy } = getCurrentInstance() as any;
        const store = useStore();
        const route = useRoute();
        const router = useRouter();
        const state: any = reactive({
            columnsAsideList: [],
            liIndex: 0,
            difference: 0,
            routeSplit: [],
        });
        // ??????????????????
        const setColumnsAsideStyle = computed(() => {
            return store.state.themeConfig.themeConfig.columnsAsideStyle;
        });
        // ??????????????????????????????
        const setColumnsAsideMove = (k: number) => {
            state.liIndex = k;
            columnsAsideActiveRef.value.style.top = `${columnsAsideOffsetTopRefs.value[k].offsetTop + state.difference}px`;
        };
        // ????????????????????????
        const onColumnsAsideMenuClick = (v: Object, k: number) => {
            setColumnsAsideMove(k);
            let { path, redirect } = v as any;
            if (redirect) router.push(redirect);
            else router.push(path);
        };
        // ????????????????????????
        const onColumnsAsideDown = (k: number) => {
            nextTick(() => {
                setColumnsAsideMove(k);
            });
        };
        // ??????/??????????????????????????????/???????????????????????????
        const setFilterRoutes = () => {
            state.columnsAsideList = filterRoutesFun(store.state.routesList.routesList);
            const resData: any = setSendChildren(route.path);
            onColumnsAsideDown(resData.item[0].k);
            proxy.mittBus.emit('setSendColumnsChildren', resData);
        };
        // ????????????????????????????????????
        const setSendChildren = (path: string) => {
            const currentPathSplit = path.split('/');
            let currentData: any = {};
            state.columnsAsideList.map((v: any, k: number) => {
                if (v.path === `/${currentPathSplit[1]}`) {
                    v['k'] = k;
                    currentData['item'] = [{ ...v }];
                    currentData['children'] = [{ ...v }];
                    if (v.children) currentData['children'] = v.children;
                }
            });
            return currentData;
        };
        // ????????????????????????
        const filterRoutesFun = (arr: Array<object>) => {
            return arr
                .filter((item: any) => !item.meta.isHide)
                .map((item: any) => {
                    item = Object.assign({}, item);
                    if (item.children) item.children = filterRoutesFun(item.children);
                    return item;
                });
        };
        // tagsView ???????????????????????????????????? columnsAsideList???????????????????????????
        const setColumnsMenuHighlight = (path: string) => {
            state.routeSplit = path.split('/');
            state.routeSplit.shift();
            const routeFirst = `/${state.routeSplit[0]}`;
            const currentSplitRoute = state.columnsAsideList.find((v: any) => v.path === routeFirst);
            // ??????????????????????????????
            setTimeout(() => {
                onColumnsAsideDown(currentSplitRoute.k);
            }, 0);
        };
        // ????????????????????????????????????????????????
        watch(store.state, (val) => {
            val.themeConfig.themeConfig.columnsAsideStyle === 'columnsRound' ? (state.difference = 3) : (state.difference = 0);
            if (val.routesList.routesList.length === state.columnsAsideList.length) return false;
            setFilterRoutes();
        });
        // ???????????????
        onMounted(() => {
            setFilterRoutes();
        });
        // ???????????????
        onBeforeRouteUpdate((to) => {
            setColumnsMenuHighlight(to.path);
            proxy.mittBus.emit('setSendColumnsChildren', setSendChildren(to.path));
        });
        return {
            columnsAsideOffsetTopRefs,
            columnsAsideActiveRef,
            onColumnsAsideDown,
            setColumnsAsideStyle,
            onColumnsAsideMenuClick,
            ...toRefs(state),
        };
    },
};
</script>

<style scoped lang="scss">
.layout-columns-aside {
    width: 64px;
    height: 100%;
    background: var(--bg-columnsMenuBar);
    ul {
        position: relative;
        li {
            color: var(--bg-columnsMenuBarColor);
            width: 100%;
            height: 50px;
            text-align: center;
            display: flex;
            cursor: pointer;
            position: relative;
            z-index: 1;
            .layout-columns-aside-li-box {
                margin: auto;
                .layout-columns-aside-li-box-title {
                    padding-top: 1px;
                }
            }
            a {
                text-decoration: none;
                color: var(--bg-columnsMenuBarColor);
            }
        }
        .layout-columns-active {
            color: #ffffff;
            transition: 0.3s ease-in-out;
        }
        .columns-round {
            background: var(--color-primary);
            color: #ffffff;
            position: absolute;
            left: 50%;
            top: 2px;
            height: 44px;
            width: 58px;
            transform: translateX(-50%);
            z-index: 0;
            transition: 0.3s ease-in-out;
            border-radius: 5px;
        }
        .columns-card {
            @extend .columns-round;
            top: 0;
            height: 50px;
            width: 100%;
            border-radius: 0;
        }
    }
}
</style>
