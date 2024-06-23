import type {AdminHealth} from "~/pkg/domain/model/health/adminHealth/AdminHealth"
import type {CommonHealth} from "~/pkg/domain/model/health/commonHealth/CommonHealth"
import type {MasterHealth} from "~/pkg/domain/model/health/masterHealth/MasterHealth"

export type HealthCheckResponse = {
    adminHealth: AdminHealth;
    commonHealth: CommonHealth;
    masterHealth: MasterHealth;
}
