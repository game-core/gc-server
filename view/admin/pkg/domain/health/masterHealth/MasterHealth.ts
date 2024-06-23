import type {MasterHealthType} from "~/pkg/domain/health/masterHealth/MasterHealthType";

export type MasterHealth = {
    "healthId": number;
    "name": string;
    "masterHealthType": MasterHealthType;
}
