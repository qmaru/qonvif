export namespace player {
	
	export class PlayParas {
	    url: string;
	    width: string;
	    height: string;
	
	    static createFrom(source: any = {}) {
	        return new PlayParas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.width = source["width"];
	        this.height = source["height"];
	    }
	}

}

export namespace ui {
	
	export class PtzAxes {
	    x: number;
	    y: number;
	
	    static createFrom(source: any = {}) {
	        return new PtzAxes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	    }
	}
	export class PtzControl {
	    name: string;
	    axes: PtzAxes;
	
	    static createFrom(source: any = {}) {
	        return new PtzControl(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.axes = this.convertValues(source["axes"], PtzAxes);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

