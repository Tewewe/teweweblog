<template>
    <div>
        <a-card>
            <h3>个人设置</h3>
            <a-form-model>
                <a-form-model-item label="作者姓名">
                  <a-input style="width: 300px"></a-input>
                </a-form-model-item>

                <a-form-model-item label="个人简介">
                    <a-input type="textarea"></a-input>
                  </a-form-model-item>

                <a-form-model-item>
                  <a-button type="danger" style="margin-right: 15px">更新</a-button>
                </a-form-model-item>
              </a-form-model>
        </a-card>
    </div>
</template>

<script>
export default {
}
</script>

<style>
</style>
