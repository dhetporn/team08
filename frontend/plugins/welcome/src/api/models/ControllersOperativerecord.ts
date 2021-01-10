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
/**
 * 
 * @export
 * @interface ControllersOperativerecord
 */
export interface ControllersOperativerecord {
    /**
     * 
     * @type {string}
     * @memberof ControllersOperativerecord
     */
    added?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersOperativerecord
     */
    examinationroom?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersOperativerecord
     */
    nurse?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersOperativerecord
     */
    operative?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersOperativerecord
     */
    tool?: number;
}

export function ControllersOperativerecordFromJSON(json: any): ControllersOperativerecord {
    return ControllersOperativerecordFromJSONTyped(json, false);
}

export function ControllersOperativerecordFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersOperativerecord {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'added': !exists(json, 'added') ? undefined : json['added'],
        'examinationroom': !exists(json, 'examinationroom') ? undefined : json['examinationroom'],
        'nurse': !exists(json, 'nurse') ? undefined : json['nurse'],
        'operative': !exists(json, 'operative') ? undefined : json['operative'],
        'tool': !exists(json, 'tool') ? undefined : json['tool'],
    };
}

export function ControllersOperativerecordToJSON(value?: ControllersOperativerecord | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added': value.added,
        'examinationroom': value.examinationroom,
        'nurse': value.nurse,
        'operative': value.operative,
        'tool': value.tool,
    };
}

