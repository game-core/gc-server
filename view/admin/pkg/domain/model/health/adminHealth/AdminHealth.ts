import type {AdminHealthType} from "~/pkg/domain/model/health/adminHealth/AdminHealthType";

export type AdminHealth = {
    "healthId": number;
    "name": string;
    "adminHealthType": AdminHealthType;
}
