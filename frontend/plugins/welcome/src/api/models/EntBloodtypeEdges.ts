/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBloodtypeEdges
 */
export interface EntBloodtypeEdges {
    /**
     * Frombloodtype holds the value of the frombloodtype edge.
     * @type {Array<EntPatient>}
     * @memberof EntBloodtypeEdges
     */
    frombloodtype?: Array<EntPatient>;
}

export function EntBloodtypeEdgesFromJSON(json: any): EntBloodtypeEdges {
    return EntBloodtypeEdgesFromJSONTyped(json, false);
}

export function EntBloodtypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBloodtypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'frombloodtype': !exists(json, 'frombloodtype') ? undefined : ((json['frombloodtype'] as Array<any>).map(EntPatientFromJSON)),
    };
}

export function EntBloodtypeEdgesToJSON(value?: EntBloodtypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'frombloodtype': value.frombloodtype === undefined ? undefined : ((value.frombloodtype as Array<any>).map(EntPatientToJSON)),
    };
}


