<template>
    <el-upload
        ref="uploadRef"
        :before-upload="beforeUpload"
        :on-success="onSuccess"
        action=""
        :http-request="httpRequest"
        :headers="headers"
        :show-file-list="showFileList"
        name="file"
        style="display: inline-block; margin-left: 2px"
    >
        <slot> </slot>
        
        <template #trigger>
            <slot name="trigger">
                <el-link icon="upload" :underline="false">上传文件夹</el-link>
            </slot>
        </template>
        
    </el-upload>
</template>
<script lang="ts">
import {ref, toRefs, reactive, onMounted, defineComponent, nextTick  } from 'vue';
import type {UploadInstance} from 'element-plus';
export default defineComponent({
    props: {
        beforeUpload: Function,
        onSuccess: Function,
        action: String,
        httpRequest: Function,
        headers:{
            type: Object,
            default: ()=> {
                return {}
            }
        },
        webkitDirectory: Boolean,
        showFileList: Boolean,
    },
    name: 'UpLoad',

    setup(props: any, { emit }) {
        const uploadRef = ref<UploadInstance>();
        const state = reactive({
            elUploadRef: null as UploadInstance | any
        })
        onMounted(async () => {
            await nextTick()
            state.elUploadRef = uploadRef.value?.$el.querySelector('input')
            state.elUploadRef.webkitdirectory = props.webkitDirectory
 
        });

        const open=()=> {
            state.elUploadRef.click();
        }

        return {
            ...toRefs(state),
            uploadRef,
            open
        }
    }
})
</script>