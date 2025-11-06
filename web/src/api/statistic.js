import service from "@/utils/request"

// 获取客户端统计数据
export const getData = (params) => {
    return service({
        url: "/user/api/v1/statistic/client",
        method: "get",
        params,
    });
};
